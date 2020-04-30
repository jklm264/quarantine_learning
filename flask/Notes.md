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

### Additional Route Decorators via Flask Plugins

   - `@login_required` (from Flask-Login): Add before any route to immediately protect it from being accessed from logged-out users. If the user is logged in, @login_required lets them in accordingly. [See more](https://flask-login.readthedocs.io/en/latest/)
    - `@cache.cached()` (from Flask-Cache): Cache routes for a set period of time, *ie:* `@cache.cached(timeout=50)`.

## Requests: Behind the scenes

- There is only a single global `request` object to every view function.
- Can hardcode routes or dynamically generate them
    - Ex: 
    ```python
    @app.route('/a/sample/<variable>/route)
    def some_view(variable):
        # some code to create variable
    ```

- `db.session` is how to connect to the db for when you want to make changes

## Templates

- For, `for loop` logic use `{%...%}` tags in html
    - Ex: 
    ``` html
        {% for contact in contacts %}
        <li>{{ contact }} </li>
        {% endfor %}
    ```
    ```python
        return render_template("index.html", contacts=['c1','c2','c3','c4','c5'])
    ```
   
### In Routing

```python
from flask import Flask, render_template, redirect, url_for
from markupsafe import escape

app = Flask(__name__)

@app.errorhandler(404)
def page_not_found(error):
    return 'This is my custom 404 page'

@app.route('/')
def index():
	return render_template('index.html')

@app.route('/1/')
def template():
    return render_template("templ8.html", message="Hello Flask!", contacts = ['c1', 'c2', 'c3', 'c4', 'c5'])

@app.route('/<drink>/')
def drinks(drink):
    return render_template(drink + '.html')

@app.route('/3/')
def redir():
	return redirect('bourbon') # will redirect to /3/bourbon which will results in 404
```

### Inheritance

- `/templates/base.html`
    
```html
<!DOCTYPE html>
<html>

<head>
    <title>Flask Template Example</title>
</head>

<body>
    <h1>Flask Template Example</h1>

    {% block content %}
    {% endblock %}

</body>
    
</html>
```

- `/templates/index.html`
```
{% extends "base.html" %}
{% block content %}
<div>
    <p>{{ message }}</p>

    <p>{{ contacts }}</p>
    <p>My Contacts:</p>
    <ul>
        {% for contact in contacts %}
        <li>{{ contact }}</li>
        {% endfor %}
    </ul>
</div>
{% endblock %}
```

- `app.py`
```python
   return render_template("index.html", message="Hello world", contacts=['c1','c2','c3','c4','c5'])
```

[Source Here](https://www.techiediaries.com/flask-tutorial-templates/)
[From the docs](https://flask.palletsprojects.com/en/1.1.x/tutorial/templates/)

---

## Installation

1. [Virtualize](howToVirtualize.md)
2. `pip install Flask`
3. `export FLASK_APP=quickstart.py`

- Make the server publically available: `$ flask run --host=0.0.0.0`
- To add [debugging](https://flask.palletsprojects.com/en/1.1.x/quickstart/#debug-mode): `export FLASK_ENV=development`
    - Will also allow [autoreload](https://stackoverflow.com/questions/16344756/auto-reloading-python-flask-app-upon-code-changes)

## Help

Use the following command if Flask says port already in use: `lsof -i -P -n | grep LISTEN`

## Secuirty

[MarkupSafe](https://markupsafe.palletsprojects.com/en/1.1.x)


## Extra Sources

- <https://hackersandslackers.com/flask-routes/>
- <https://flask.palletsprojects.com>
