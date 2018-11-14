#include<stdio.h>
#include<math.h>
#define EPS .00000001
double func_f(double, double);
int main(int argc, char **argv){
	double x=1.0,y=1.0;
	double h=0.1,dx=1.0,xmax=10.0;
	double ddx=0.0,k1,k2,k3,k4;
	int i;

	printf("h\t X\t\tY\n");
	for(i=0;i<=10;i++){
		do{
			if(x>=ddx-EPS){
				ddx+=dx;
				if(x>=1.9 && x<=2.1)
				printf("%8.4lf %8.4lf %8.4lf\n",h,x,y);
			}
			k1=func_f(x,y);
			k2=func_f(x+h/2.0,y+h*k1*h/2.0);
			k3=func_f(x+h/2.0,y+h*k1*h/2.0);
			k4=func_f(x+h,y+h+k3*h);

			y+=(h/6.0)*(k1+2.0*k2+2.0*k3+k4);
			x+=h;
		}while(x<=xmax);
		h=h/2.0;
		x=1.0;
		y=1.0;
		ddx=0.0;
	}
	return 0;
}
double func_f(double x,double y){
	return 3.0*y/(1+x);
}
