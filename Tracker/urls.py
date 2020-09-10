from django.urls import path
from . import views

urlpatterns = [
    path('', views.home, name='tracker-home'),
    path('search/', views.search, name='tracker-search'),
    path('result/', views.result, name='tracker-result'),

]