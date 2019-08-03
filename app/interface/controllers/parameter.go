package controllers

import "strconv"

func getLimit(ctx Context) (limit int) {
	limit, err := strconv.Atoi(ctx.Query("limit"))
	if err != nil {
		limit = 10
	}
	return
}

func getOffset(ctx Context) (offset int) {
	offset, err := strconv.Atoi(ctx.Query("offset"))
	if err != nil {
		offset = 0
	}
	return
}

func getLectureID(ctx Context) (lectureID int64) {
	id, err := strconv.Atoi(ctx.Param("lecture_id"))
	if err != nil {
		id = 0
	}
	lectureID = int64(id)
	return
}

func getHelpID(ctx Context) (helpID int64) {
	id, err := strconv.Atoi(ctx.Param("help_id"))
	if err != nil {
		id = 0
	}
	helpID = int64(id)
	return
}

func getCommentID(ctx Context) (commentID int64) {
	id, err := strconv.Atoi(ctx.Param("comment_id"))
	if err != nil {
		id = 0
	}
	commentID = int64(id)
	return
}
