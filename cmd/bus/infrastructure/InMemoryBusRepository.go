package infrastructure

import "back-transportes-fisi/cmd/bus/domain"

type InMemoryBusRepository struct {
	buses []*domain.Bus
}

func NewInMemoryBusRepository() *InMemoryBusRepository {
	return &InMemoryBusRepository{buses: make([]*domain.Bus, 0)}
}

func (this *InMemoryBusRepository) Create(bus *domain.Bus) error {
	this.buses = append(this.buses, bus)
	return nil
}

func (this *InMemoryBusRepository) GetAll() ([]*domain.Bus, error) {
	return this.buses, nil
}

func (this *InMemoryBusRepository) GetById(id int) (*domain.Bus, error) {
	for _, bus := range this.buses {
		if bus.IdBus == id {
			return bus, nil
		}
	}

	return nil, nil
}

func (this *InMemoryBusRepository) Update(bus *domain.Bus) error {
	for i, b := range this.buses {
		if b.IdBus == bus.IdBus {
			this.buses[i] = bus
			return nil
		}
	}

	return nil
}

func (this *InMemoryBusRepository) Delete(id int) error {
	var newBuses []*domain.Bus
	for _, bus := range this.buses {
		if bus.IdBus != id {
			newBuses = append(newBuses, bus)
		}
	}
	this.buses = newBuses

	return nil
}
