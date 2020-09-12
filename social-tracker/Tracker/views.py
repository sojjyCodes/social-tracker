from django.shortcuts import render
from django.http import HttpResponse
import requests
from pprint import pprint

# Create your views here.

title = [{
    "title": "social-tracker"
}]

api_key = [{
   " username": "sojjyCodes",
}]

def home(request):
    home_title = {
        'title': title
    }

    return render(request, 'social-tracker/index.html', home_title)

def search(request):
    search_title = {
        'title': title
}
    return render(request, 'social-tracker/search.html', search_title)

def result(request):
    result_title = {
       'title': title
    }
    authentication = {
        'api_key': api_key
    }

    return render(request, 'social-tracker/results.html', authentication)

# def api(request):
#     # github username
#     username = "sojjyCodes"
#     # url to request
#     url = f"https://api.github.com/users/{username}"
#     # make the request and return the json
#     user_data = requests.get(url).json()

#     # pretty print JSON data
#     pprint(user_data)