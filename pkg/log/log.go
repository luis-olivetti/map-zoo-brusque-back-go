package log

import "go.uber.org/zap"

type Logger struct {
	*zap.Logger
}

func NewLogger() *Logger {
	// Configuração do logger Zap
	config := zap.NewProductionConfig()
	// Desabilita a stack trace em produção
	config.DisableStacktrace = true
	// Cria um logger Zap
	logger, err := config.Build()
	if err != nil {
		panic(err)
	}
	// Retorna uma instância do Logger personalizado
	return &Logger{logger}
}
