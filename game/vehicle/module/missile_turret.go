package module

import (
	"fmt"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/lucas-s-work/warships/game/projectile"
	"github.com/lucas-s-work/warships/game/world"
)

type MissileTurret struct {
	*SmallTurret
}

func CreateMissileTurret(w world.World, pos mgl32.Vec2, parent world.Entity) *MissileTurret {
	return &MissileTurret{
		SmallTurret: CreateSmallTurret(w, pos, parent),
	}
}

func (s *MissileTurret) OnFire(w world.World, event world.KeyInputEvent) {
	if s.ammoCount == 0 {
		return
	}

	fmt.Println(s.Position())
	target := event.MousePos.Add(event.CameraPos)
	if b := projectile.CreateMissile(w, mgl32.Vec3{target.X(), target.Y(), 0}, s.Position()); b != nil {
		s.ammoCount -= 1
		w.AttachEntity(b)
	}
}
