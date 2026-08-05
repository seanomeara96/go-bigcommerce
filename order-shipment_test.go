package bigcommerce

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
)

func TestGetOrderShipments(t *testing.T) {
	require.NoError(t, godotenv.Load(".env"))

	client := NewClient(os.Getenv("BF_STORE_HASH"), os.Getenv("BF_XAUTHTOKEN"), nil, nil)
	require.NotNil(t, client)

	shipments, _, err := client.V2.GetOrderShipments(9419527, OrderShipmentQueryParams{})
	require.NoError(t, err)
	require.NotZero(t, len(shipments))
	t.Logf("%+v\n", shipments)
}
