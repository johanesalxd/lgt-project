package cli_test

import "testing"

func assertPlayerWin(t testing.TB, store *StubPlayerStore, winner string) {
	t.Helper()

	if len(store.winCalls) != 1 {
		t.Fatalf("got %d calls to RecordWin want %d", len(store.winCalls), 1)
	}

	if store.winCalls[0] != winner {
		t.Errorf("incorrect winner got %q want %q", store.winCalls[0], winner)
	}
}

func assertScheduledAlert(t testing.TB, got, want scheduledAlert) {
	amountGot := got.amount
	if amountGot != want.amount {
		t.Errorf("got %d want %d", amountGot, want.amount)
	}

	scheduleGot := got.at
	if scheduleGot != want.at {
		t.Errorf("got %v want %v", scheduleGot, want.at)
	}
}
