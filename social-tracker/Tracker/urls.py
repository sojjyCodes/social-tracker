from django.urls import path, include
from django.contrib.auth.models import User
from rest_framework import routers, serializers, viewsets

class UserSerializer(serializers.HyperlinkedModelSerializer):
    class Meta:
        model = User
        fields = ['info']

class UserViewSet(viewsets.ModelViewSet):
    queryset = User.objects.all()
    serializer_class = UserSerializer

router = routers.DefaultRouter()
router.register(r'users', UserViewSet)

urlpatterns = [
    path('', views.home, name='tracker-home'),
    path('api/', include(router.urls)),
    path('search/', views.search, name='tracker-search'),
    path('result/', views.result, name='tracker-result'),

]
