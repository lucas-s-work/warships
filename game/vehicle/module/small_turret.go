package module

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/lucas-s-work/warships/game/projectile"
	"github.com/lucas-s-work/warships/game/world"
)

const (
	maxAmmo   = 10
	maxHealth = 10

	smallTurretTexture = "./textures/weapons/WeaponDestroyerStandardGun.png"
	smallTurretWidth   = 15
	smallTurretHeight  = 26
	smallTurretHealth  = 50
)

var (
	smallTurretAimCenter = mgl32.Vec2{7, 8}
)

type SmallTurret struct {
	*BaseModule
	health    int
	ammoCount int
}

func CreateSmallTurret(w world.World, pos mgl32.Vec2, parent world.Entity) *SmallTurret {
	module := CreateBaseModule(
		w,
		pos,
		parent,
		smallTurretTexture,
		mgl32.Vec2{
			smallTurretWidth,
			smallTurretHeight,
		},
		smallTurretAimCenter,
		smallTurretHealth,
	)

	return &SmallTurret{
		BaseModule: module,
		ammoCount:  maxAmmo,
		health:     maxHealth,
	}
}

func (s *SmallTurret) AmmoCount() int {
	return s.ammoCount
}

func (s *SmallTurret) Health() int {
	return s.health
}

func (s *SmallTurret) Reload(int) {
}

func (s *SmallTurret) OnFire(w world.World, event world.KeyInputEvent) {
	if s.ammoCount == 0 {
		return
	}

	target := event.MousePos.Add(event.CameraPos)
	if b := projectile.CreateBullet(w, target, s.Position()); b != nil {
		s.ammoCount -= 1
		w.AttachEntity(b)
	}
}

func (s *SmallTurret) Sprite() mgl32.Vec4 {
	return mgl32.Vec4{
		0,
		0,
		smallTurretWidth,
		smallTurretHeight,
	}
}

func (s *SmallTurret) Texture() string {
	return smallTurretTexture
}
