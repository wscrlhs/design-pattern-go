// Package factorymethod
// 简单工厂方法，通常没有子类。 但当从一个简单工厂中抽取出子类后， 它看上去就会更像经典的工厂方法模式了。
// class UserFactory {
// 	public static function create($type) {
// 	    switch ($type) {
// 		case 'user': return new User();
// 		case 'customer': return new Customer();
// 		case 'admin': return new Admin();
// 		default:
// 		    throw new Exception('传递的用户类型错误。');
// 	    }
// 	}
//     }

// 工厂方法需要用到继承，其在父类中提供一个创建对象的方法， 允许子类决定实例化对象的类型。
//abstract class Department {
// 	public abstract function createEmployee($id);

// 	public function fire($id) {
// 	    $employee = $this->createEmployee($id);
// 	    $employee->paySalary();
// 	    $employee->dismiss();
// 	}
//     }

//     class ITDepartment extends Department {
// 	public function createEmployee($id) {
// 	    return new Programmer($id);
// 	}
//     }

//	    class AccountingDepartment extends Department {
//		public function createEmployee($id) {
//		    return new Accountant($id);
//		}
//	    }
package factorymethod

import "fmt"

// IGun 产品接口
type IGun interface {
	SetName(name string)
	SetPower(power int)
	GetName() string
	GetPower() int
	Show() string
}

// Gun 具体产品
type Gun struct {
	name  string
	power int
}

func (g *Gun) SetName(name string) {
	g.name = name
}

func (g *Gun) GetName() string {
	return g.name
}

func (g *Gun) SetPower(power int) {
	g.power = power
}

func (g *Gun) GetPower() int {
	return g.power
}

func (g *Gun) Show() string {
	return "this is " + g.name
}

// Ak47 具体产品
type Ak47 struct {
	Gun
}

func newAk47() IGun {
	return &Ak47{
		Gun: Gun{
			name:  "AK47 gun",
			power: 4,
		},
	}
}

// musket 具体产品
type musket struct {
	Gun
}

func newMusket() IGun {
	return &musket{
		Gun: Gun{
			name:  "Musket gun",
			power: 1,
		},
	}
}

// GetGun 包含大量条件语句的构建方法， 可根据方法的参数来选择对何种产品进行初始化并将其返回。
func GetGun(gunType string) (IGun, error) {
	if gunType == "ak47" {
		return newAk47(), nil
	}
	if gunType == "musket" {
		return newMusket(), nil
	}
	return nil, fmt.Errorf("Wrong gun type passed")
}
