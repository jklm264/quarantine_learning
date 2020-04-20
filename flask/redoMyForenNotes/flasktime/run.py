from flask import Flask, render_template, url_for, request, jsonify

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

@app.route('/api', methods=['GET'])
def api_get():
    print('Queried from: ', request.headers.get('User-Agent'))
    if 'girl' in request.args:
        print('Its a girl')
        return 'Its a girl!'
    else:
        print('Its a boy, thank you!')
        return 'Its a boy, thank you!'
    # return jsonify(books)

@app.route('/api/add', methods=['POST'])
def api_post():
    print(request.json.get('boy'))
    return jsonify('You just posted a boy')

@app.errorhandler(404)
def pageNotFound(error):
    return render_template("404.html")
app.register_error_handler(404, pageNotFound)

# Some easy test data
# books = [
#     {'id': 0,
#      'title': 'A Fire Upon the Deep',
#      'author': 'Vernor Vinge',
#      'first_sentence': 'The coldsleep itself was dreamless.',
#      'year_published': '1992'},
#     {'id': 1,
#      'title': 'The Ones Who Walk Away From Omelas',
#      'author': 'Ursula K. Le Guin',
#      'first_sentence': 'With a clamor of bells that set the swallows soaring, the Festival of Summer came to the city Omelas, bright-towered by the sea.',
#      'published': '1973'},
#     {'id': 2,
#      'title': 'Dhalgren',
#      'author': 'Samuel R. Delany',
#      'first_sentence': 'to wound the autumnal city.',
#      'published': '1975'}
# ]
