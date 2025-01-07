package setup

import (
	"keyvalue-api/constants"

	"github.com/supabase-community/supabase-go"
)

var SupabaseClient *supabase.Client

func InitSupabase() error {

	client, err := supabase.NewClient(constants.SUPABASE_URL, constants.SUPABASE_ANON_KEY, nil)
	if err != nil {
		return err
	}

	SupabaseClient = client
	return nil
}
