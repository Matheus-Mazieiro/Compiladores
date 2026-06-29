#include <ilcplex/ilocplex.h>
#include <vector>
#include <iostream>

ILOSTLBEGIN

int main(){
	IloEnv env;
	
	try {
		IloModel model(env);
		
		
		IloCplex cplex(model);
		
		if(cplex.solve()){
			std::cout << "Objective = " << cplex.getObjValue() << std::endl;
		}
	}
	catch(IloException& e){
		std::cerr << e << std::endl;
	}
	
	env.end();
	return 0;
}
