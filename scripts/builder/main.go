package main

import (
	"encoding/json"
	"os"
	"time"

	"github.com/117503445/goutils"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/rs/zerolog/log"
)

type VersionInfo struct {
	BuiltAt time.Time
	Commit  string
	Tag     string
	IsDirty bool
}

func main() {
	goutils.InitZeroLog()

	if err := os.Chdir("../../"); err != nil {
		log.Fatal().Err(err).Msg("change dir failed")
	}

	versionInfo := VersionInfo{
		BuiltAt: time.Now(),
	}

	repo, err := git.PlainOpen(".")
	if err != nil {
		log.Fatal().Err(err).Msg("open repo failed")
	}

	// get is dirty
	wt, err := repo.Worktree()
	if err != nil {
		log.Fatal().Err(err).Msg("get worktree failed")
	}
	status, err := wt.Status()
	if err != nil {
		log.Fatal().Err(err).Msg("get status failed")
	}
	versionInfo.IsDirty = !status.IsClean()

	// get latest commit
	ref, err := repo.Head()
	if err != nil {
		log.Fatal().Err(err).Msg("get head failed")
	}
	commit, err := repo.CommitObject(ref.Hash())
	if err != nil {
		log.Fatal().Err(err).Msg("get commit failed")
	}
	versionInfo.Commit = commit.Hash.String()

	// get the tags of the commit
	tags, err := repo.Tags()
	if err != nil {
		log.Fatal().Err(err).Msg("get tags failed")
	}
	tags.ForEach(func(ref *plumbing.Reference) error {
		if ref.Hash() == commit.Hash {
			if versionInfo.Tag != "" {
				log.Fatal().Msg("multiple tags found")
			}
			versionInfo.Tag = ref.Name().Short()
		}
		return nil
	})

	log.Info().Interface("versionInfo", versionInfo).Msg("version info")

	// write to version.json
	js, err := json.Marshal(versionInfo)
	if err != nil {
		log.Fatal().Err(err).Msg("marshal version info failed")
	}
	if err := os.WriteFile("version.json", js, 0644); err != nil {
		log.Fatal().Err(err).Msg("write version info failed")
	}
}
