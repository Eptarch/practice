package erratum

func Use(opener ResourceOpener, input string) (e error) {
	resource, err := opener()
	if err != nil {
		switch err.(type) {
		    case TransientError: return Use(opener, input)
		    default: return err
		}
	}
    defer resource.Close()
	defer func() {
		if recovery := recover(); recovery != nil {
			switch recovery.(type) {
			case FrobError:
				e = recovery.(FrobError).inner
				resource.Defrob(recovery.(FrobError).defrobTag)
			case error:
				e = recovery.(error)
			}
		}
	}()
	resource.Frob(input)
	return
}
