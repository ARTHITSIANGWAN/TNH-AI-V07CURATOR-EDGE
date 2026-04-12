package main

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/syumai/workers"
	"github.com/syumai/workers/adapter"
	"github.com/syumai/workers/cloudflare/kv"
	
	// 📍 พิกัด Import ต้องเปลี่ยนตามชื่อ Repo ใหม่ของบอส!
	"github.com/ARTHITSIANGWAN/thitnueahub-post-the-curator-V7/pkg/core" 
)

func main() {
	app := fiber.New(fiber.Config{
		CaseSensitive: false,
	})

	// 🎭 The Curator: Content Matrix Ignite
	app.Post("/v7/curate", func(c *fiber.Ctx) error {
		// Get from LayerPool (Zero Garbage)
		layer := core.LayerPool.Get().(*core.ContentLayer)
		defer core.LayerPool.Put(layer)

		if err := c.BodyParser(layer); err != nil {
			return c.Status(400).JSON(fiber.Map{"status": "error", "msg": "invalid matrix"})
		}

		// Save to KV (3d180d...98)
		ctx := context.Background()
		_ = kv.Namespace("KV").Put(ctx, "last_curated_hook", []byte(layer.Hook), nil)

		return c.JSON(fiber.Map{
			"status": "curated",
			"worker": "tnh-curator",
			"hook":   layer.Hook,
			"engine": "zero-garbage-2000%",
		})
	})

	// 🥄 Plug and Play (No Listen!)
	workers.Serve(adapter.FiberApp(app))
}

