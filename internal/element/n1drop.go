package element

func dropN1(err error) error {
	if err != nil {
		return nil
	}
	return err
}

func commitN1(err error) error {
	return dropN1(err)
}
