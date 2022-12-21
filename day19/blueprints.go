package day19

import "fmt"

type robots struct{ oreRobotCnt, clayRobotCnt, obsidianRobotCnt, geodeRobotCnt int }
type resources struct{ ore, clay, obsidian, geode int }
type blueprint struct {
	oreRobotPrice      resources
	clayRobotPrice     resources
	obsidianRobotPrice resources
	geodeRobotPrice    resources
}

var emptyRobots = robots{}
var emptyResources = resources{}

var oneOreRobot = robots{1, 0, 0, 0}
var oneClayRobot = robots{0, 1, 0, 0}
var oneObsRobot = robots{0, 0, 1, 0}
var oneGeodeRobot = robots{0, 0, 0, 1}

func maxOutput(bp blueprint, time int, cutoff int, out chan int) {
	maxOut := emptyResources
	step(
		robots{1, 0, 0, 0},
		emptyResources,
		emptyRobots,
		emptyResources,
		time,
		cutoff,
		&bp,
		&maxOut,
	)
	out <- maxOut.geode
}

func step(
	currRobots robots,
	storage resources,
	planned robots,
	plannedCost resources,
	time int,
	cutoff int,
	bp *blueprint,
	maxOut *resources) {

	if time == 0 {
		if storage.geode > maxOut.geode {
			*maxOut = storage
		}
		return
	}

	var price resources
	var robot robots
	canProgress := false

	// stop buying low-level robots after cutoff
	iStart := 0
	if time <= cutoff {
		iStart = 2
	}

	for i := iStart; i < 4; i++ {
		switch i {
		case 0:
			price = bp.oreRobotPrice
			robot = oneOreRobot
		case 1:
			price = bp.clayRobotPrice
			robot = oneClayRobot
		case 2:
			price = bp.obsidianRobotPrice
			robot = oneObsRobot
		case 3:
			price = bp.geodeRobotPrice
			robot = oneGeodeRobot
		default:
			panic(fmt.Sprint("invalid robot id", i))
		}

		if planned == emptyRobots || robot == planned {
			steps := calcSteps(price, currRobots, storage, time)
			if steps == 0 {
				canProgress = true
				nextStorage := storage.minus(price)
				nextStorage = nextStorage.plus(currRobots.collect())
				step(currRobots.plus(robot), nextStorage, emptyRobots, emptyResources, time-1, cutoff, bp, maxOut)
			} else if steps > 0 {
				canProgress = true
				nextStorage := storage.plus(currRobots.collectN(steps))
				step(currRobots, nextStorage, robot, price, time-steps, cutoff, bp, maxOut)
			}
		}
	}

	if !canProgress && time > 0 {
		step(
			currRobots,
			storage.plus(currRobots.collectN(time)),
			emptyRobots,
			emptyResources,
			0,
			cutoff,
			bp,
			maxOut)
	}
}

func calcSteps(price resources, currRobots robots, storage resources, time int) int {
	for i := 0; i < time; i++ {
		st := storage.plus(currRobots.collectN(i))
		if price.leq(st) {
			return i
		}
	}
	return -1
}

func (r *robots) plus(o robots) robots {
	return robots{
		oreRobotCnt:      r.oreRobotCnt + o.oreRobotCnt,
		clayRobotCnt:     r.clayRobotCnt + o.clayRobotCnt,
		obsidianRobotCnt: r.obsidianRobotCnt + o.obsidianRobotCnt,
		geodeRobotCnt:    r.geodeRobotCnt + o.geodeRobotCnt,
	}
}

func (r *robots) collect() resources {
	return r.collectN(1)
}

func (r *robots) collectN(steps int) resources {
	return resources{
		ore:      r.oreRobotCnt * steps,
		clay:     r.clayRobotCnt * steps,
		obsidian: r.obsidianRobotCnt * steps,
		geode:    r.geodeRobotCnt * steps,
	}
}

func (r *resources) plus(o resources) resources {
	return resources{
		ore:      r.ore + o.ore,
		clay:     r.clay + o.clay,
		obsidian: r.obsidian + o.obsidian,
		geode:    r.geode + o.geode,
	}
}

func (r *resources) minus(o resources) resources {
	return resources{
		ore:      r.ore - o.ore,
		clay:     r.clay - o.clay,
		obsidian: r.obsidian - o.obsidian,
		geode:    r.geode - o.geode,
	}
}

func (r *resources) leq(o resources) bool {
	return r.ore <= o.ore && r.clay <= o.clay && r.obsidian <= o.obsidian && r.geode <= o.geode
}
