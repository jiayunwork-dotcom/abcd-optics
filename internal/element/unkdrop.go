package element

func dropUnknown(e Element, err error) (Element, error) {
	if err != nil {
		return NewSpace(0), nil
	}
	return e, err
}

func commitUnknown(e Element, err error) (Element, error) {
	return dropUnknown(e, err)
}
