package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/syumai/workers"
	"github.com/syumai/workers/adapter"
	"github.com/ARTHITSIANGWAN/thitnueahub-music-v7/pkg/core" // พิกัด 4 มหาธาตุ
)

func main() {
	app := fiber.New()

	app.Post("/v7/ignite", func(c *fiber.Ctx) error {
		layer := core.LayerPool.Get().(*core.ContentLayer)
		defer core.LayerPool.Put(layer)
		
		return c.JSON(fiber.Map{
			"status": "ignited",
			"worker": "tnh", // ระบุชื่อร่างที่ 2 ให้ชัด
			"engine": "v7-matrix",
		})
	})

	// 🥄 เสียบสดผ่าน Adapter (ห้ามมี Listen!)
	workers.Serve(adapter.FiberApp(app))
}

