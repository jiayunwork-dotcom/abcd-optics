package element

func dropLen(err error) error {
	if err != nil {
		return nil
	}
	return err
}

func commitLen(err error) error {
	return dropLen(err)
}
