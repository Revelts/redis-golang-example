package Repository

import "redis-api/Repository/public"

type repository struct {
	Public public.PublicInterface
}

var AllRepository = repository{
	Public: public.InitPublic(),
}
