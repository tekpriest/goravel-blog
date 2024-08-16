package observers

import (
	"fmt"

	"github.com/goravel/framework/contracts/database/orm"
)

type UserModelObserver struct{}

func (u *UserModelObserver) Retrieved(event orm.Event) error {
	return nil
}

func (u *UserModelObserver) Creating(event orm.Event) error {
	fmt.Println("here all the way Creating")
	fmt.Println(event)
	return nil
}

func (u *UserModelObserver) Created(event orm.Event) error {
	fmt.Println("here all the way Created")
	fmt.Println(event)
	return nil
}

func (u *UserModelObserver) Updating(event orm.Event) error {
	return nil
}

func (u *UserModelObserver) Updated(event orm.Event) error {
	return nil
}

func (u *UserModelObserver) Saving(event orm.Event) error {
	fmt.Println(event)
	return nil
}

func (u *UserModelObserver) Saved(event orm.Event) error {
	fmt.Println(event)
	return nil
}

func (u *UserModelObserver) Deleting(event orm.Event) error {
	return nil
}

func (u *UserModelObserver) Deleted(event orm.Event) error {
	return nil
}

func (u *UserModelObserver) ForceDeleting(event orm.Event) error {
	return nil
}

func (u *UserModelObserver) ForceDeleted(event orm.Event) error {
	return nil
}
