package component

type ServiceComponent struct {
	controlService IControlService
}

type IServiceComponent interface {
	GetControlService() *IControlService
}

func NewServiceComponent(controlService IControlService) *ServiceComponent {
	return &ServiceComponent{controlService: controlService}
}

func (s ServiceComponent) GetControlService() *IControlService {
	return &s.controlService
}
