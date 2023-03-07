package main

func main() {
	// Hardcoded account number and AES encoded password
	accountNumber := "352457987298754320975239087209857349028"
	AESPassword := "a87d7f9eue9f9d7d87fd8783f439rf398f8934r82f948fg934r498gh493384y9839834y98y4983y4983g98w4yf9833gw49g83984f94hf98h49f834hf9834hf9834hhf983hhf934hf39488fh9348hf9348hf9348fh9384yr98232hf98hh1389f39384wgf383949fh98fh982h982hfh298hf298hf98ch298fh298hf9298h39823"

	// Assign to struct variable values of account number and password
	var newAccount = AccountEntry{AccNumber: accountNumber, Password: AESPassword}

	// Connecting to DB
	ConnectDB(newAccount)
}
