# API
This is a coding-time document mainly about APIs in this project.

## user
### /api/signup
JSON

	{
		"email" : "example@example.com",
		"username" : "Bob",
		"password" : "password",
		"gender" : 0,
		"birthday" : "2001-01-01"
	}

Return Code:

    U100 : Success!
    U101 : Cannot connect to the database.
    U102 : SQL statement error
    U103 : Unable to get DB handle
    U104 : Execution error
    
### /api/checkusername/
JSON

	{
		"username" : "Bob"
	}

Return Code:

    C100-1 : There is a same username
    C100-0 : There is no same username
    C101 : Cannot connect to the database.
    C102 : SQL statement error
    C103 : SQL execution error
    
### /api/checkemail/
JSON

	CheckEmail is a function that handles
	{
		"email" : "example@example.com"
	}

Return Code:

    C100-1 : There is a same email
    C100-0 : There is no same email
    C101 : Cannot connect to the database.
    C102 : SQL statement error
    C103 : SQL execution error
    
### /api/login/
JSON

	{
		"username" : "Bob",
		"password" : "password"
	}

Return Code:

    C100-1 : OK
    C100-0 : Wrong password
    C101 : Cannot connect to the database.
    C102 : SQL statement error
    C103 : SQL execution error
    C104 : No user
    U200 : Can not login again!

### /api/logout/
JSON

    None

Return Code:

    U100 : OK
    U200 : Not logged in yet!

## forum
### /api/getPosts
JSON

    {
        path : "/"
    }