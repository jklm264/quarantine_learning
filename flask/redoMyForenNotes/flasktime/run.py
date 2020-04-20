from flask import Flask, render_template, url_for

app = Flask(__name__)
app.config['TEMPLATES_AUTO_RELOAD'] = True
app.run(debug=True, host='0.0.0.0')
app.jinja_env.auto_reload = True

@app.route("/", methods=['GET'])
def index():
    return render_template("index.html", title='My-OFFLINE-Forensics-Notes')

@app.route("/Notes/<page>", methods=['GET'])
def notes(page):
    if page.endswith('.html'):
        page = page[:-5]

    if page.lower() == 'windows':
        return render_template("windows.html")
    elif page.lower() == 'apple':
        return render_template("apple.html")
    elif page.lower() == 'cloud':
        return render_template("cloud.html")
    elif page.lower() == 'osint':
        return render_template("osint.html")
    else:
        return pageNotFound(404)

@app.errorhandler(404)
def pageNotFound(error):
    return render_template("404.html")
app.register_error_handler(404, pageNotFound)
