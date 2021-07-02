package config

import (
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	Load()

	if Items().ProjectEnv != Dev {
		t.Errorf("Load() = %v, want %v", Items().ProjectEnv, Dev)
	}

	os.Setenv("PROJECT_ENV", Prod)
	Load()
	if Items().ProjectEnv != Prod {
		t.Errorf("Load() = %v, want %v", Items().ProjectEnv, Prod)
	}

	os.Setenv("MYSQL_PORT", "1234")
	Load()
	if Items().Mysql.Port != "1234" {
		t.Errorf("Load() = %v, want %v", Items().ProjectEnv, "1234")
	}

	// cannot assign to Items().APIVersion
	// Items().APIVersion = "v1.2.3"
}
