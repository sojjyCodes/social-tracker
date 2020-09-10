from django.shortcuts import render
from django.http import HttpResponse
import requests
from pprint import pprint

# Create your views here.


def home(request):
    return render(request, 'social-tracker/index.html')

def search(request):
    return render(request, 'social-tracker/search.html')

def result(request):
    return render(request, 'social-tracker/results.html')

def api(request):
    # github username
    username = "sojjyCodes"
    # url to request
    url = f"https://api.github.com/users/{username}"
    # make the request and return the json
    user_data = requests.get(url).json()

    # pretty print JSON data
    pprint(user_data)
    return render