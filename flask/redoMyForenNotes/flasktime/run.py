from flask import Flask, render_template, url_for, request, jsonify, abort

app = Flask(__name__)
app.config['TEMPLATES_AUTO_RELOAD'] = True
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

@app.route('/api', methods=['GET'])
def api_get():
    if 'topic' in request.args:
        print('A topic was provided')
        return 'Thank you for providing a topic to research. Too bad the db hasnt been created yet.'
    else:
        print('A topic was not provided.')
        return 'No topic provided. Remember to include a topic!'

# Put is used for updating resources while post should create new ones.
@app.route('/api/add', methods=['POST'])
def api_post():
    print('\nTopic is> ', request.json.get('topic'))
    print('Queried from: ', request.headers.get('User-Agent'))
    return jsonify('You just **posted** a topic!')

@app.errorhandler(404)
def pageNotFound(error):
    return render_template("404.html")
app.register_error_handler(404, pageNotFound)

if __name__ == "__main__":
    app.run(debug=True, host='...')
