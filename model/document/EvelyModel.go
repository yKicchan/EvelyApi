package document

type evelyModel struct {
	Event       *EventModel
	User        *UserModel
	PendingUser *PendingUserModel
}

type EvelyModel func(*evelyModel)

func (this EvelyModel) Make() evelyModel {
	model := evelyModel{}
	this(&model)
	return model
}

func Event(e *EventModel) EvelyModel {
	return func(this *evelyModel) {
		this.Event = e
	}
}

func User(u *UserModel) EvelyModel {
	return func(this *evelyModel) {
		this.User = u
	}
}

func PendingUser(pu *PendingUserModel) EvelyModel {
	return func(this *evelyModel) {
		this.PendingUser = pu
	}
}
