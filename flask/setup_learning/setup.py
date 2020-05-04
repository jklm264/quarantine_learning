import os
from setuptools import setup

# Utility function to read the README file.
# Used for the long_description.  It's nice, because now 1) we have a top level
# README file and 2) it's easier to type in the README file than to put a raw
# string in below ...
def read(fname):
    return open(os.path.join(os.path.dirname(__file__), fname)).read()

setup(
    name = "packagemeup",
    version = "0.0.4",
    author = "My name",
    author_email = "myemail@no.com",
    description = ("A short little intro project"),
    license = "BSD",
    keywords = "example documentation tutorial",
    url = "https://github.com/jklm264",
    packages=['packagemeup', 'tests'],
    long_description=read('README.md'),
    long_description_content_type="text/markdown",
    classifiers=[
        "Development Status :: 3 - Alpha",
        "Topic :: Utilities",
        'Topic :: Scientific/Engineering',
        "License :: OSI Approved :: BSD License",
    ],
)


# Src: https://pythonhosted.org/an_example_pypi_project/setuptools.html
