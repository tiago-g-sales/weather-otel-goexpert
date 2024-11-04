package web

import (
	"encoding/json"
	"io"
	"net/http"

	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/tiago-g-sales/weather-otel-goexpert/configs"
	"github.com/tiago-g-sales/weather-otel-goexpert/internal/model"


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
	BackgroundColor    string
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

	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	
	
	carrier := propagation.HeaderCarrier(r.Header)
	ctx := r.Context()
	ctx = otel.GetTextMapPropagator().Extract(ctx, carrier)

	ctx, spanInicial := h.TemplateData.OTELTracer.Start(ctx, "SPAN_INICIAL"+h.TemplateData.RequestNameOTEL)
	time.Sleep(time.Second)
	spanInicial.End()

	ctx, span := h.TemplateData.OTELTracer.Start(ctx, "Chama externa"+h.TemplateData.RequestNameOTEL)
	defer span.End()

	time.Sleep(time.Millisecond * h.TemplateData.ResponseTime)

	if configs.ExternalCallURL != "" {
		var req *http.Request
		var err error

		cepParam := r.URL.Query().Get(QUERY_PARAMETER)
		if cepParam == ""{
			http.Error(w,"Invalid Parameter CEP", http.StatusUnprocessableEntity )
			return
		}
		
		if configs.ExternalCallMethod == "GET" {
			req, err = http.NewRequestWithContext(ctx, "GET", configs.ExternalCallURL, nil)
		} else if configs.ExternalCallMethod == "POST" {
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
		
		h := http.Client{}

		q := req.URL.Query() 
		q.Add("cep", cepParam)
		req.URL.RawQuery = q.Encode()	
		resp, err := h.Do(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
	
		temp := model.Temperatura{}

		err= json.Unmarshal(body, &temp)
		if err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(temp)

	

	

	}


}
