# gocheckenv

Simple utility to check if certain environment variables exist. Returns len 0 slice of strings if valid or len N slice of strings that includes all missing env variables.

## USAGE:

Create a slice of strings with the required environment variables you want to locate. Pass them to gocheckenv.CheckEnv() and check the length of the response slice.
