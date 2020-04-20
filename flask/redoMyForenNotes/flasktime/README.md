# Notes for this flask app

## Creating Env

1. `$virtualenv -p /usr/bin/python3 ./path_of_virtenv`
2. `$source ./path_of_virtenv/bin/activate`
3. `$export FLASK_APP=run.py`
4. `$flask run` # If changing python (not html), will need to rerun this cmd
5. `$deactivate`

## Creating APIs

- [Better API POST source](https://blog.miguelgrinberg.com/post/designing-a-restful-api-with-python-and-flask)
- https://programminghistorian.org/en/lessons/creating-apis-with-python-and-flask

### Troubleshooting

- Use the following command if Flask says port already in use: `lsof -i -P -n | grep LISTEN then kill <pid>`

## Running the app

- Flask is very name specific so make sure templates dir is in the root directory, else you will have to overwrite the path- which is a hassel.This includes images (serve static files), flask will look for the `/static` file
- Also, do not name your `{root}.py` app.py (especially if you have a subfolder called `app/`). The preference is to name it `run.py` and then make sure to run `$export FLASK_APP='run.py'`.
- To get app to reload on change, add the following to `run.app`:

    ```python
    app.config['TEMPLATES_AUTO_RELOAD'] = True
    app.run(debug=True, host='0.0.0.0')
    app.jinja_env.auto_reload = True
    ```

## Bootstrap 4 navbar tips

- <https://getbootstrap.com/docs/4.0/components/navbar/>

## Jinja Templating

- <https://code.tutsplus.com/tutorials/templating-with-jinja2-in-flask-essentials--cms-25571>
- Note, that if the child template doesn’t define the footer block, the value from the parent template is used instead.
