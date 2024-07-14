package cron

func initArticleData() {
	// # ┌───────────── min (0 - 59)
	// # │┌────────────── hour (0 - 23)
	// # ││┌─────────────── day of month (1 - 31)
	// # │││┌──────────────── month (1 - 12)
	// # ││││┌───────────────── day of week (0 - 6) (0 to 6 are Sunday to
	// # │││││         Saturday, or use names; 7 is also Sunday)
	// # │││││
	// # │││││
	// # * * * * * command to execute
	entryID, err := crontab.AddFunc("00 */5 * * * ?", DeferFunc(refreshArticleData))
	if err != nil {
		panic("add cron task error. taskName: refreshArticleData")
	}
	logger.Info("add cron task success. taskName: refreshArticleData, entryID:", entryID)
}

func refreshArticleData() {
	logger.Info("刷新数据完成")
}
