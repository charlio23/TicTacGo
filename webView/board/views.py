from django.shortcuts import render
from django.http import HttpResponse, HttpResponseRedirect
from django.shortcuts import get_object_or_404, render
# Create your views here.

def index(request):
	return render(request, 'board/board.html')