# Enron email-indexer

## This application uses:
- Zincsearch for indexing
- Go for parsing files and as an api to connect zincsearch and the frontend
- chi as an api router
- Vue 3 as the frontend framework
- Tailwind for css

## Steps for running the application:
- Download and unzip the enron database
- Run the file-reading file, make sure the paths are correctly set up. (root := "../enron_mail_20110402")
- Download zincsearch and start it, the password and username used in the code is admin:Complexpass#123
- Using zincsearch's ui, add a new template using the enron_mail_index.txt file
- Create a parsed-files directory inside of the file-reading directory
- Run the bulk-insert file
- Run npm install inside of the email-indexer folder to download all node dependencies
- Run npm run dev
- The application shoould be good to go on http://localhost:5173/
