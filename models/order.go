package models

import "gorm.io/gorm"

type NodeMeta struct {
	gorm.Model
	NodeRef int    `json:"-,omitempty"`
	Node    Node   `gorm:"foreignKey:NodeRef;" json:"node"`
	Icon    string `json:"icon"`
}

type Node struct {
	gorm.Model

	NodeID          string       `json:"nodeId"`
	SequenceID      int          `json:"sequenceId"`
	Released        bool         `json:"released"`
	Actions         []Action     `gorm:"many2many:node_actions;" json:"actions"`
	NodeDescription string       `json:"nodeDescription"`
	NodePositionID  int          `json:"-,omitempty"` // Not  in the vda55 struct, simply a field for GORM
	NodePosition    NodePosition `gorm:"foreignKey:NodePositionID;" json:"nodePosition"`
}

type Action struct {
	gorm.Model

	ActionID          string            `json:"actionId"`
	ActionType        string            `json:"actionType"`
	BlockingType      string            `json:"blockingType"`
	ActionDescription string            `json:"actionDescription"`
	ActionParameters  []ActionParameter `gorm:"many2many:action_actionparameters;" json:"actionParameters"`
}

type ActionParameter struct {
	gorm.Model

	Key   string  `json:"key"`
	Value float64 `json:"value"`
}

type NodePosition struct {
	gorm.Model

	X                     float64 `json:"x"`
	Y                     float64 `json:"y"`
	MapID                 string  `json:"mapId"`
	Theta                 float64 `json:"theta"`
	AllowedDeviationXY    float64 `json:"allowedDeviationXY"`
	AllowedDeviationTheta float64 `json:"allowedDeviationTheta"`
	MapDescription        string  `json:"mapDescription"`
}

type Edge struct {
	gorm.Model

	EdgeID           string     `json:"edgeId"`
	SequenceID       int        `json:"sequenceId"`
	Released         bool       `json:"released"`
	StartNodeID      string     `json:"startNodeId"`
	EndNodeID        string     `json:"endNodeId"`
	Actions          []Action   `gorm:"many2many:edge_actions;" json:"actions"`
	EdgeDescription  string     `json:"edgeDescription"`
	MaxSpeed         float64    `json:"maxSpeed"`
	MaxHeight        float64    `json:"maxHeight"`
	MinHeight        float64    `json:"minHeight"`
	Orientation      float64    `json:"orientation"`
	OrientationType  string     `json:"orientationType"`
	Direction        string     `json:"direction"`
	RotationAllowed  bool       `json:"rotationAllowed"`
	MaxRotationSpeed float64    `json:"maxRotationSpeed"`
	Length           float64    `json:"length"`
	TrajectoryID     int        `json:"-,omitempty"` // Not  in the vda55 struct, simply a field for GORM
	Trajectory       Trajectory `gorm:"foreignKey:TrajectoryID;" json:"trajectory"`
	CorridorID       int        `json:"-,omitempty"` // Not  in the vda55 struct, simply a field for GORM
	Corridor         Corridor   `gorm:"foreignKey:CorridorID;" json:"corridor"`
}

type Trajectory struct {
	gorm.Model
	Degree int `json:"degree"`
	// KnotVector    []float64      `json:"knotVector"`
	ControlPoints []ControlPoint `gorm:"many2many:trajectory_controlpoints;" json:"controlPoints"`
}

type ControlPoint struct {
	gorm.Model
	ControlPointID int
	X              float64 `json:"x"`
	Y              float64 `json:"y"`
	Weight         float64 `json:"weight"`
}

type Corridor struct {
	gorm.Model
	LeftWidth        float64 `json:"leftWidth"`
	RightWidth       float64 `json:"rightWidth"`
	CorridorRefPoint string  `json:"corridorRefPoint"`
}

// Not VDA 5050, but needed for storing before sending to a robot.
type OrderTemplateDetails struct {
	gorm.Model
	Name            string        `json:"name"`
	OrderTemplateId int           `json:"-,omitempty"`
	Order           OrderTemplate `gorm:"foreignKey:OrderTemplateId;" json:"order"`
	NodeIds         []string      `gorm:"-" json:"nodeIds"` //This is only for processing when receving from rest
}

// Having the order template be the exact same struct as the Order
// is for us to be able to have two actual database tables with Gorm.
// This means that we can create a template and use that for when creating
// an actual order. It is important to keep this struct the same as the Order
// struct, as we can simply clone the template into the non-template,
// when we want to send an order.
type OrderTemplate struct {
	gorm.Model
	HeaderID      int    `json:"headerId"`
	Timestamp     string `json:"timestamp"`
	Version       string `json:"version"`
	Manufacturer  string `json:"manufacturer"`
	SerialNumber  string `json:"serialNumber"`
	OrderID       string `json:"orderId"`
	OrderUpdateID int    `json:"orderUpdateId"`
	Nodes         []Node `gorm:"many2many:order_template_nodes;" json:"nodes"`
	Edges         []Edge `gorm:"many2many:order_template_edges;" json:"edges"`
	ZoneSetID     string `json:"zoneSetId"`
}

// Since we have the order template for building the actual templating for the orders
// this means that whenever we save an actual order, we will save the order in its
// entirity, filled out with nodes, edges and their actions.
type Order struct {
	gorm.Model
	HeaderID      int    `json:"headerId"`
	Timestamp     string `json:"timestamp"`
	Version       string `json:"version"`
	Manufacturer  string `json:"manufacturer"`
	SerialNumber  string `json:"serialNumber"`
	OrderID       string `json:"orderId"`
	OrderUpdateID int    `json:"orderUpdateId"`
	Nodes         []Node `gorm:"many2many:order_nodes;" json:"nodes"`
	Edges         []Edge `gorm:"many2many:order_edges;" json:"edges"`
	ZoneSetID     string `json:"zoneSetId"`
}
