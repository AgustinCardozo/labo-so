package log

import (
	"fmt"
	"io"
	"log/slog"
	"os"
)

func InitLogger(logPath string, logLevel string) {
	//Creamos el archivo "modulo".log en modo escritura, si ocurre algún error finalizamos con panic.
	logFile, err := os.OpenFile(logPath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)

	if err != nil {
		panic(err)
	}

	// Usa io.MultiWriter para escribir a múltiples destinos: consola y archivo.
	multiWriter := io.MultiWriter(os.Stdout, logFile)

	// Convertir el Log_Level del config al tipo slog.Level.
	level, err := convertStringToLogLevel(logLevel)

	// Crear un nuevo manejador de registros (handler) para colocar.
	handler := slog.NewTextHandler(multiWriter, &slog.HandlerOptions{
		Level: level, // Toma el valor del level_log que se define arriba con convertirLogLevelDeConfig.
	})

	// Configura slog para que use el manejador que creamos anteriormente.
	slog.SetDefault(slog.New(handler))

	// Escribimos en el log el warning que obtenemos por no setear el log_level
	if err != nil {
		slog.Warn(err.Error())
	}

	slog.Debug("Se ha configurado correctamente el logger y el archivo de configuración. ")
}

// El objetivo de esta función es seterar dinámicamente el nivel de log que deseamos tener en el sistema.
func convertStringToLogLevel(levelStr string) (slog.Level, error) {
	switch levelStr {
	case "DEBUG":
		return slog.LevelDebug, nil
	case "INFO":
		return slog.LevelInfo, nil
	case "WARN":
		return slog.LevelWarn, nil
	case "ERROR":
		return slog.LevelError, nil
	default:
		return slog.LevelInfo, fmt.Errorf("No existe %s, se coloca INFO por defecto. ", levelStr)
	}
}
