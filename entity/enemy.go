package entity

import (
/*
	"fmt"
	. "github.com/daniel840829/gameServer/msg"
	"github.com/daniel840829/gameServer/physic"
	//p "github.com/golang/protobuf/proto"
	//"github.com/gazed/vu/math/lin"
	//. "github.com/daniel840829/gameServer/uuid"
	"github.com/golang/protobuf/proto"
	"github.com/ianremmler/ode"
	"github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"os"
	"reflect"
	"sync"
	//"time"
*/
)

type Enemy struct {
	Player
}

func (e *Enemy) PhysicUpdate() {
	//AI
}