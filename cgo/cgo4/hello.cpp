#include <stdio.h>
#include <stdlib.h>

extern "C" void MyPrint() {
 printf("Hello world from C\n");
}



#include <iostream>
#include <vector>
#include <string>

using namespace std;

extern "C" void MyPrintCPP() {
    vector<string> v;
    v.push_back("Hello world");
    v.push_back("from C++");
    cout << v[0] << " " << v[1] << endl;
}
