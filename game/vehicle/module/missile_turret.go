package module

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/lucas-s-work/warships/game/projectile"
	"github.com/lucas-s-work/warships/game/world"
)

const (
	MissileTurretWidth   = 13
	MissileturretHeight  = 13
	MissileTurretTexture = "./textures/weapons/WeaponMissileLauncher.png"
)

type MissileTurret struct {
	*BaseModule
	ammoCount int
}

func CreateMissileTurret(w world.World, pos mgl32.Vec2, parent world.Entity) *MissileTurret {
	module := CreateBaseModule(
		w,
		pos,
		parent,
		MissileTurretTexture,
		mgl32.Vec2{
			MissileTurretWidth,
			MissileturretHeight,
		},
		mgl32.Vec2{7, 7},
		10,
	)

	module.SetOffset(mgl32.Vec2{0, -7})

	return &MissileTurret{
		BaseModule: module,
		ammoCount:  6,
	}
}

func (s *MissileTurret) AmmoCount() int {
	return s.ammoCount
}

func (s *MissileTurret) Health() int {
	return s.health
}

func (s *MissileTurret) Reload(int) {
}

func (s *MissileTurret) OnFire(w world.World, event world.KeyInputEvent) {
	if s.ammoCount == 0 {
		return
	}

	target := event.MousePos.Add(event.CameraPos)
	if b := projectile.CreateMissile(w, mgl32.Vec3{target.X(), target.Y(), 0}, s.Position()); b != nil {
		s.ammoCount -= 1
		w.AttachEntity(b)
	}
}

func (*MissileTurret) SetTarget(mgl32.Vec2) {}

func (s *MissileTurret) Sprite() mgl32.Vec4 {
	return mgl32.Vec4{
		0,
		0,
		MissileTurretWidth,
		MissileturretHeight,
	}
}

func (s *MissileTurret) Texture() string {
	return MissileTurretTexture
}
