#include "hello.h"
#include <iostream>

using namespace std;

int main()
{
	cout << "C++ says: calling HelloWorld" << endl;
	std::string str = HelloWorld();
	cout << str << endl;
	return 0;
}

