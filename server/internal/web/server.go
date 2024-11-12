package web

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/viper"
	"github.com/tiago-g-sales/weather-otel-goexpert/internal/model"
	"github.com/tiago-g-sales/weather-otel-goexpert/pkg"
	"github.com/valyala/fastjson"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)



type Webserver struct {
	TemplateData *TemplateData
}

// NewServer creates a new server instance
func NewServer(templateData *TemplateData) *Webserver {
	return &Webserver{
		TemplateData: templateData,
	}
}

// createServer creates a new server instance with go chi router
func (we *Webserver) CreateServer() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)
	router.Use(middleware.Timeout(60 * time.Second))
	// promhttp
	router.Handle("/metrics", promhttp.Handler())
	router.Get("/", we.HandleRequest)
	return router
}

type TemplateData struct {
	Title              string
	ResponseTime       time.Duration
	ExternalCallMethod string
	ExternalCallURL    string
	Content            string
	RequestNameOTEL    string
	OTELTracer         trace.Tracer
}

const(
	INVALID_ZIP_CODE = "invalid zipcode"
	CAN_NOT_FIND_ZIPCODE = "can not find zipcode"
	QUERY_PARAMETER = "cep"
	LEN_ZIP_CODE = 8
)

func (h *Webserver) HandleRequest(w http.ResponseWriter, r *http.Request) {

	carrier := propagation.HeaderCarrier(r.Header)
	ctx := r.Context()
	ctx = otel.GetTextMapPropagator().Extract(ctx, carrier)

	ctx, spanInicial := h.TemplateData.OTELTracer.Start(ctx, "SPAN_INICIAL"+h.TemplateData.RequestNameOTEL)
	spanInicial.End()

	ctx, span := h.TemplateData.OTELTracer.Start(ctx, "Chamada externa CEP "+h.TemplateData.RequestNameOTEL)
	defer span.End()

	if h.TemplateData.ExternalCallURL != "" {
		var req *http.Request
		var err error
		
		if r.URL.Path != "/"{
			w.WriteHeader(http.StatusNotFound)
			return
		}

		cepParam := r.URL.Query().Get(QUERY_PARAMETER)
		if cepParam == ""{
			http.Error(w,"Invalid Parameter CEP", http.StatusUnprocessableEntity )
			return
		}

		if len(cepParam) > LEN_ZIP_CODE || len(cepParam) < LEN_ZIP_CODE  {
			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(INVALID_ZIP_CODE)
			return		
		}

		if h.TemplateData.ExternalCallMethod == "GET" {
			req, err = http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("http://viacep.com.br/ws/%s/json/",cepParam), nil)
		} else if h.TemplateData.ExternalCallMethod == "POST" {
			req, err = http.NewRequestWithContext(ctx, "POST", h.TemplateData.ExternalCallURL, nil)
		} else {
			http.Error(w, "Invalid ExternalCallMethod", http.StatusInternalServerError)
			return
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(req.Header))
		http_client := http.Client{}

		q := req.URL.Query() 
		req.URL.RawQuery = q.Encode()	
		resp, err := http_client.Do(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
	
		var c model.ViaCEP
		err= json.Unmarshal(body, &c)
		if err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		temp := model.Temperatura{}

		ctx, span = h.TemplateData.OTELTracer.Start(ctx, "Chamada externa Temperatura "+h.TemplateData.RequestNameOTEL)
		defer span.End()

		req, err = http.NewRequestWithContext(ctx, "GET", "https://api.weatherapi.com/v1/current.json", nil)
		if err != nil {
			return 
		}
		q = req.URL.Query() 
		q.Add("q", c.Localidade)
		q.Add("key", viper.GetString("API_KEY"))
		req.URL.RawQuery = q.Encode()		
	
		resp, err = http_client.Do(req)
		if err != nil {
			return 
		}
		defer resp.Body.Close()

		var p fastjson.Parser

		body, err = io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		v, err := p.ParseBytes(body)
		if err != nil {
			panic(err)
		}

		json.Unmarshal([]byte(v.GetObject("current").String()), &temp)
		
		temp.Temp_K, _ = pkg.ConvertTemp(temp.Temp_C)
		temp.City = c.Localidade
		
		

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(temp)

	

	

	}


}
