# Dinning Philosophers problem

The dining philosophers problem is a computer science problem that illustrates how concurrent processes behave. It's often used to teach about synchronization issues and how to resolve them

## Contributors

* Adeyemi Alexandra Adeola - 230591018
* Ekojikoko Oghenemudia Peter - 230591070
* Keripe Olufemi Benua - 230591272
* Ridwan Pedro Olawale - 210591152
* Abiola OLamilekan Marcus - 230591352
* Awodunmila Olaoluwa Eniola - 240591383
* Idehen Elijah Osaigbovo - 240591396
* Amanze Samuel Ikechukwu - 230591042
* Davies Omotola Ibrahim - 230591064
* Pelujo Shukurat Mojisola - 240519431
* Ogundipe Odebiyi Oluwaseyifunmi - 230591122
* Ayedun Inioluwa - 230591050
* Agboola Mayowa - 230591363
* Adebayo Mojeed Alade - 230591006
* Gbadebo Esther Adebimpe - 230591294
* Bakare Mubashir Oladimeji - 240591408
* Sanni Azeezat Abolanle - 230591340
* Cole Toluwani Joshua - 230591373
* Suraju Abdulrahmon - 230591256
* Akin-Akala Oluwaponmile Oluwatimilehin - 230591030
* Ochoja Favour - 240591408
* Bakare Victor Oluwasetuntunfunmi - 230591191
* Okebule Olumide Michael - 230591138
* Adekunle Quadri Kanyisola - 230591191
* Ajorin Emmanuel - 230591287
* Adeoti Michael Oluwapamilerin - 230591012
* Adebayo Feyishayo - 230591308
* Akinola Abdul-Rahman Akinwunmi - 230591122
* Ogazie Chinoso Miracle - 230591118

## Key Project Areas

1. **Philosopher Definition:**
   ```bash
   type Philosopher struct {
        id    int
        left  *sync.Mutex
        right *sync.Mutex
   }
Here the Philospher object is created and each of the attribute is defined.

The Philospher has an Id for identification, a Left fork and also a Right fork which is needed for the philosopher to eat.

2. **Philosopher Attributues:**
   ```go
  func (p *Philosopher) eat() {
        fmt.Printf("Philosopher %d is eating\n", p.id)
        time.Sleep(time.Duration(rand.Intn(3)) * time.Second) // Simulate eating
        fmt.Printf("Philosopher %d finished eating\n", p.id)

    }
  
  func (p *Philosopher) think() {

        fmt.Printf("Philosopher %d is thinking\n", p.id)
        time.Sleep(time.Duration(rand.Intn(3)) * time.Second) // Simulate thinking
  }

  func (p *Philosopher) dine() {
        for {
                p.think()

                fmt.Printf("Philosopher %d is trying to eat\n", p.id)

                // Acquire locks -  Important: Order matters to prevent deadlock!
                p.left.Lock()
                fmt.Printf("Philosopher %d acquired left fork\n", p.id)
                p.right.Lock() // If this blocks, the left is held!
                fmt.Printf("Philosopher %d acquired right fork\n", p.id)

                p.eat()

                // Release locks - Order doesn't matter here, but good practice to reverse.
                p.right.Unlock()
                fmt.Printf("Philosopher %d released right fork\n", p.id)
                p.left.Unlock()
                fmt.Printf("Philosopher %d released left fork\n", p.id)


        }
  }

 Every Philosopher as defined  in  this code has attributes like eat, think and dine
  

