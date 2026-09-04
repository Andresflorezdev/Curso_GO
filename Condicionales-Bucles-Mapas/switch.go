package main

import (
	"errors"
	"fmt"
	"time"
)

type UserTier string

const (
	TierFree     UserTier = "FREE"
	TierPro      UserTier = "PRO"
	TierEnterprise UserTier = "ENTERPRISE"
)

type Request struct {
	UserID    string
	Tier      UserTier
	Cost      int
	Config    any // Interfaz vacía con parámetros avanzados
}

type EnterpriseConfig struct {
	DedicatedNode string
	MaxPriority    bool
}

type RateLimiter struct{}

func (rl *RateLimiter) Evaluate(req Request) (time.Duration, error) {
	// 1. Switch condicional con inicialización en línea
	// Evalúa condiciones de bloqueo inmediato según consumo y horarios
	switch currentHour := time.Now().Hour(); {
	case req.Cost <= 0:
		return 0, errors.New("el costo de la petición debe ser mayor a 0")
	case req.Cost > 100 && req.Tier != TierEnterprise:
		return 0, errors.New("cuota excedida: solo Enterprise admite peticiones de alto consumo")
	case currentHour >= 2 && currentHour <= 4 && req.Tier == TierFree:
		// Ventana de mantenimiento nocturno para usuarios Free
		return 0, errors.New("servicio Free en mantenimiento programado")
	}

	// 2. Switch con fallthrough para definir capacidades en cascada
	var cooldown time.Duration
	var features []string

	switch req.Tier {
	case TierEnterprise:
		features = append(features, "SLA Prioritario")
		fallthrough // Pasa al siguiente case para heredar capacidades Pro
	case TierPro:
		features = append(features, "Soporte 24/7", "Webhooks")
		cooldown = 100 * time.Millisecond
	case TierFree:
		features = append(features, "Acceso Básico")
		cooldown = 2 * time.Second
	default:
		return 0, fmt.Errorf("nivel de usuario desconocido: %s", req.Tier)
	}

	fmt.Printf("[PERMISOS] Usuario: %s | Nivel: %s | Capacidades: %v\n", req.UserID, req.Tier, features)

	// 3. Type Switch para procesar configuraciones específicas del payload
	switch cfg := req.Config.(type) {
	case EnterpriseConfig:
		if cfg.MaxPriority {
			cooldown = 0 // Bypass de cooldown para prioridad máxima
			fmt.Printf(" -> Ruteando hacia nodo dedicado: %s\n", cfg.DedicatedNode)
		}
	case map[string]string:
		if val, exists := cfg["custom_cooldown"]; exists {
			fmt.Printf(" -> Cooldown personalizado detectado: %s\n", val)
		}
	case nil:
		// Configuración vacía permitida
	default:
		return 0, fmt.Errorf("tipo de configuración no soportado: %T", cfg)
	}

	return cooldown, nil
}

func condicionesmain() {
	limiter := &RateLimiter{}

	// Ejemplo: Petición de usuario Enterprise aprovechando la herencia de permisos
	req := Request{
		UserID: "usr_ent_99",
		Tier:   TierEnterprise,
		Cost:   50,
		Config: EnterpriseConfig{
			DedicatedNode: "us-east-1a",
			MaxPriority:    true,
		},
	}

	cooldown, err := limiter.Evaluate(req)
	if err != nil {
		fmt.Printf("Petición rechazada: %v\n", err)
		return
	}

	fmt.Printf("Petición aceptada. Tiempo de espera requerido: %v\n", cooldown)
}
