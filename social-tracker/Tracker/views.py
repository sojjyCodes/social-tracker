from django.shortcuts import render
from django.http import HttpResponse

# Create your views here.

api = [{
   "title": "social-tracker",
   "username": "sojjyCodes",
   "followers": 50,
   "repository": 2,
   "contribution": 346,
}]

def home(request):
    home_title = {
        'api': api
    }

    return render(request, 'social-tracker/index.html', home_title)

def search(request):
    search_title = {
        'api': api
}
    return render(request, 'social-tracker/search.html', search_title)

def result(request):
    result_title = {
        'api': api
}
    return render(request, 'social-tracker/results.html', result_title)
# def api(request):
#      github username
#      username = "sojjyCodes"
#      url to request
#      url = f"https://api.github.com/users/{username}"
#      # make the request and return the json
#     user_data = requests.get(url).json()

# #     # pretty print JSON data
# #     pprint(user_data)