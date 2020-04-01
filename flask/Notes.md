# Flask Notes

- [See Install and quickstart](#Installation)
- [Troubleshooting](#Help)

## Different Python Web Frameworks

- Flask
- Pyramid - aimed at taking small web apps into big web apps. [Read more](https://trypyramid.com)
- Django

## This Proj

Inital structure from [opensource.com](https://opensource.com/article/18/4/flask)

## Setup

- Python `setuptools` library is a collection of utils that allow devs to easily build and distribute python packages, especially those with many dependencies. [More](https://setuptools.readthedocs.io/en/latest/setuptools.html#basic-use)
- Will use `flask-sqlalchemy` for the SQL dbms
- Will use `psycopg2` for the *shim* between the app and the dbms
- The `__init__.py` file allows you to import from your project as if it were an installed package.
- The `app.py` file will be the application's root. This is where all the *Flask* application goodness will go, and you'll create an environment variable that points to that file.

## `app.py`

- The central config object for the entire application
- Used to set up app's *extended* functionality like db access, crypto, ect.
- **Routes** - `view` functions used to access the application

For user data to persist in web apps, we use databases.

## Database

- When using flask-sqlalchemy, need a DBAPI-2 standard DBMS intermediary. Here we use PostgreSQL
- In a given web application, you may want to be able to express relationships between objects. In the To-Do List example, users own multiple tasks, and each task is owned by only one user. This is an example of a "many-to-one" relationship, also known as a foreign key relationship, where the tasks are the "many" and the user owning those tasks is the "one."
    - Use `db.relationship`
    - `back_populates` keyword updates the User model when the task gets a user as an owner.
- Last to initalize, you just need something to recognize what tables are to be created and some code to create them (or drop them should the need arise

## Objects

- a "model" refers to the data representation of some real or conceptual object
    - For example, if you're building an application for a car dealership, you may define a Car model that encapsulates all of a car's attributes and behavior.

- `db.Model` object and define attributes with `db.Colum` instance

## Views and Routes

- A **view** is the functionality that runs when a specific access/end point (aka **route**) is "hit".
- Routes are URLs/paths
- "With Flask, a function is marked as a view when it is decorated by app.route." Learn more about decorators [here](https://flask.palletsprojects.com/en/1.1.x/patterns/viewdecorators/)
- `app.route` is the app's central configurations map for routing
    - Makes sense it's located iun `app.py` (the app's "central config object")

- Recommended to use Python dictionaries to encapsulate the HTTP requests/responses

- Note: Flask is finiky with URIs. Best practice is to include `/` at and of path for redirection. More on this [here](https://flask.palletsprojects.com/en/1.1.x/quickstart/#unique-urls-redirection-behavior)

## Requests: Behind the scenes

- There is only a single global `request` object to every view function.
- Can hardcode routes or dynamically generate them
    - Ex: 
    ```
    @app.route('/a/sample/<variable>/route)
    def some_view(variable):
        # some code to create variable
    ```

- `db.session` is how to connect to the db for when you want to make changes

---

## Installation

1. `pip install Flask`
2. Then [virtualize](howto_virtualize.md)
3. [Docs](https://flask.palletsprojects.com/en/1.1.x/quickstart/#quickstart)
4. `export FLASK_APP=quickstart.py`

- Make the server publically available: `$ flask run --host=0.0.0.0`
- To add [debugging](https://flask.palletsprojects.com/en/1.1.x/quickstart/#debug-mode): `export FLASK_ENV=development`
    - Will also allow [autoreload](https://stackoverflow.com/questions/16344756/auto-reloading-python-flask-app-upon-code-changes)

## Help

Use the following command if Flask says port already in use: `lsof -i -P -n | grep LISTEN`

## Secuirty

[MarkupSafe](https://markupsafe.palletsprojects.com/en/1.1.x)