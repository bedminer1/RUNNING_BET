// See https://kit.svelte.dev/docs/types#app
// for information about these interfaces
// and what to do when importing types
declare namespace App {
	interface Locals {
		pb: import('pocketbase').default
	}
	// interface PageData {}
	// interface Error {}
	// interface Platform {}
}

interface WeeklyRecord {
	weekID: number
	// put these 3 as optional so that placeholder in form fields can show up
	myScore?: number 
	herScore?: number
	neededScore?: number
	winForMe: boolean
	myPoints: number
	herPoints: number
}

interface RawWeeklyRecord {
	myScore?: number
	herScore?: number
	neededScore?: number
	winForMe: boolean
}

interface DataSet {
	label: string
	data: number[]
	borderColor: string
	backgroundColor: string
}