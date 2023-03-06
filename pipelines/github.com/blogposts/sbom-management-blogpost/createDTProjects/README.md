# create DT Projects

The files in this directory use github's command line client and Dependency Track's API to create a Dependency Track project for every repository in a Github organization.

You need to have Go and Github's command line client installed in order to run these successfully.

The script in `create-projects.sh` takes care of running `main.go`.
You need to have exported as environment variables

* `DT_API_KEY`  : the api key for the target Dependency Track
* `DT_URL` : the url of the target Dependency Track

Then you can execute the script as such

```bash
DT_API_KEY=<> DT_URL=<> ./create-projects.sh
```