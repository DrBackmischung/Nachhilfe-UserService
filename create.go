package model
import (
	"log"
	"sync"
	users "../lib"
	"github.com/nu7hatch/gouuid"
)
func CreateSkill(e Skill, label string, c chan error, mutex *sync.Mutex) {

// beginning of the critical section
	mutex.Lock()
	u, err := uuid.NewV4()
	result, err := Session.Run(`
	CREATE (n:INCHARGE {
		id:$id,
		name:$name,
		level:$level)<-[:`+label+`]-(a) `,
		map[string]interface{}{
			"id":           e.GetField(label, "id"),
			"name": 		e.GetField(label, "name"),
			"level":        e.GetField(label, "level")
		}
	)
	if err != nil {
		c <- err
		return
	}
// critical section ends
	mutex.Unlock()
	if err = result.Err(); err != nil {
		c <- err
		return
	}
	log.Printf("Created %s node", label)
	c <- nil
	return
}