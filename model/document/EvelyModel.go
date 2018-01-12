package document

type evelyModel struct {
	event       *EventModel
	user        *UserModel
	pendingUser *PendingUserModel
}

type EvelyModel func(*evelyModel)

func (this EvelyModel) make() evelyModel {
	model := evelyModel{}
	this(&model)
	return model
}

func (this EvelyModel) GetEvent() *EventModel {
	return this.make().event
}

func (this EvelyModel) GetUser() *UserModel {
	return this.make().user
}

func (this EvelyModel) GetPendingUser() *PendingUserModel {
	return this.make().pendingUser
}

func Event(e *EventModel) EvelyModel {
	return func(this *evelyModel) {
		this.event = e
	}
}

func User(u *UserModel) EvelyModel {
	return func(this *evelyModel) {
		this.user = u
	}
}

func PendingUser(pu *PendingUserModel) EvelyModel {
	return func(this *evelyModel) {
		this.pendingUser = pu
	}
}
