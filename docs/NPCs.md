# Non-Player Characters

NPCs on the Mars site have a single shared AI. This AI is driven by a few different attributes of the NPC.

- Skill
	- Specialization score
- Desires
- Emotional Health
	- Motivation type
- Hunger
- Energy / Rest...

## Tasks

The AI picks which tasks the NPC does. It picks one task at a time. The task it picks is the task the NPCs desire has ranked highest that it can do based on the tasks difficulty rating.

The tasks they do drive their emotional health. Their emotional health while doing a task impacts how they learn. Being stressed eats away at skill points in the active category. Being in the growth zone results in accelerated learning. Being in the comfort zone results in minimal to no learning.

Each task has a difficulty rating from 0 to X. Task difficulty starts at its default level and is then reduced by the NPCs specific skills related to the task and the skill level in the skills type. And then the difficulty is augmented by the NPCs emotional health and NPCs desire level to see the task done. An NPC cannot do a task of more than 10 difficulty.

By default an NPC will only be aware of tasks use skills at their level or one level higher. It is possible to assign tasks that require a higher skill level to an NPC but NPCs can only do tasks from their skill sets level or lower.

There are two types of tasks. Personal tasks like eating, sleeping, relaxing and learning and site wide tasks like "mine 100 rocks" or "build a building at X,Y".

All site wide tasks have a skill (eg: harvest, build a building at X,Y) a complete quantity an available quantity and a target.

"Harvest 100 carrots" -> "skill quantity target"
"Build a building at X,Y" -> "skill quantity skill (at target)"

The available quantity is the quantity of the target that is currently on site. Different skills calculate this differently. For "harvest" the calculation would be the number of the item fully grown on plants. For "build a building" the calculation would be based on the resources needed to build a building.

The skill and the quantity will determine the number of people who can be working on it at a time.

Example:
"Build building" will allow 2 people per 1 building
"Harvest" will allow 1 person per 20 of the item available.

## Emotional Health

Emotional health is a scale from 0 to 9. This scale is divided into their health windows: comfort, growth and stress. Each window will never have a width of less than 1 or more than 6 and windows do not overlap. 0 is always the start of comfort. 9 is always the end of stress.

An NPCs emotional health is driven by their active task and how much their desires are being met.

Tasks with a difficulty rating of less than 5 drive the NPC towards 0. Tasks with a difficulty rating of more than 5 drive the NPC towards 9. The longer a task takes the more impact the task's difficulty has on their emotional health.

Desires impact emotional health daily. Desires that have lasted less than 3 days have no impact on emotional health. Desires then have lated more than 3 days drive emotional health towards 9. When a desire that is more than one day old is accomplished it drives emotional health towards 0.

Moving emotional health towards zero when in the comfort window, grows the comfort window by reducing stress and then growth.

Moving towards nine in the comfort window has no effect on the health windows.

Moving towards zero in the growth window grows growth by reducing stress.

Moving towards nine in the growth window grows stress by reducing comfort.

Moving towards nine in the stress window grows stress by reducing growth and then comfort.

Moving towards zero in the stress window grows stress by reducing comfort and then growth.

## Desires




## Skill Sets

There are X skill sets. Each skill set has Y levels. Each level has different skills it unlocks. You can learn a skill at your level or lower by doing it. This will increase the difficulty. You can learn a skill one above your level by reading about it.

Each skill level represents a skill difficulty change of 4 points. In each level skills have a difficulty bonus of 0 to 2. Doing a skill you don't know has a difficulty bound of 2.

So a skills difficulty equals (SKILL_LEVEL * 4) + DIFFICULTY_BONUS + KNOWLEDGE_BOUNS.

So a skill you know at SKILL_LEVEL 2 that has a DIFFICULTY_BONUS of 1 has a difficulty rating of 9. This is its default rating. You then have a capacity rating which reduces this rating. Your capacity starts at (SKILL_LEVEL - 1 * 4). So at skill level 1 (which is the lowest level) you have a base capacity of 0. This means its technically possible to do the SKILL_LEVEL 2 task mentioned above, but just barley.

Being in the stress zone decreases capacity by 2.
Being in growth zone increases capacity by 1.
Being in the comfort zone does nothing to capacity.

This means that doing a SKILL with a difficulty rating of 9 when you have a base of 0 and are in the stress zone is not possible. Though it would be possible do a task with a difficulty rating of 11 if you were in the growth zone.

## Task groups