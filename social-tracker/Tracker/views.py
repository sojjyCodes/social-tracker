from django.shortcuts import render
from django.http import HttpResponse

# Create your views here.
def home(request):
    return render(request, 'social-tracker/index.html')

def search(request):
    return render(request, 'social-tracker/search.html')

def result(request):
    return render(request, 'social-tracker/results.html')