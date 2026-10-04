package config

import (
	"github.com/joho/godotenv"
	"github.com/rpstvs/mcp-go/utils"
)

type Config struct {
	TargetDirectoryRepo string
}

func Load() *Config {

	godotenv.Load(".env")

	repoDir := utils.GetEnvString("target_repo_dir", "fallback")

	return &Config{
		TargetDirectoryRepo: repoDir,
	}
}
