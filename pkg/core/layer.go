package core

import "sync"

// ContentLayer: 4 มหาธาตุที่ห้ามขาด (Hook, Body, Caption, Thumbnail)
type ContentLayer struct {
	Hook      string // 3 วิแรกต้องหยุดโลก
	Body      string // คุณค่าเน้นๆ (สุนทรีย์/ความรู้)
	Caption   string // ตัวดันยอดแชร์
	Thumbnail string // สรุปความสงสัยให้คนคลิก
}

// Zero Garbage: Pool สำหรับ ContentLayer ลดการจอง Memory
var LayerPool = sync.Pool{
	New: func() interface{} {
		return new(ContentLayer)
	},
}

