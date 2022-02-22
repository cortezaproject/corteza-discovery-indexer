package app

import (
	"context"
	"github.com/cortezaproject/corteza-discovery-indexer/indexer"
	"github.com/cortezaproject/corteza-discovery-indexer/searcher"
)

const (
	bootLevelWaiting = iota
	bootLevelSetup
	bootLevelStoreInitialized
	bootLevelProvisioned
	bootLevelServicesInitialized
	bootLevelActivated
)

// Setup configures all required services
func (app *CortezaDiscoveryApp) Setup() (err error) {
	app.lvl = bootLevelSetup
	return nil
}

// InitStore initializes open search store and runs upgrade procedures
func (app *CortezaDiscoveryApp) InitStore(ctx context.Context) (err error) {
	if app.lvl >= bootLevelStoreInitialized {
		// Is store already initialised?
		return nil
	} else if err = app.Setup(); err != nil {
		// Initialize previous level
		return err
	}

	app.lvl = bootLevelStoreInitialized
	return nil
}

// Provision instance with configuration and settings
// by importing preset configurations and running autodiscovery procedures
func (app *CortezaDiscoveryApp) Provision(ctx context.Context) (err error) {
	if app.lvl >= bootLevelProvisioned {
		return
	}

	if err = app.InitStore(ctx); err != nil {
		return err
	}

	app.lvl = bootLevelProvisioned
	return
}

// InitServices initializes all services used
func (app *CortezaDiscoveryApp) InitServices(ctx context.Context) (err error) {
	if app.lvl >= bootLevelServicesInitialized {
		return nil
	}

	if err = app.Provision(ctx); err != nil {
		return err
	}

	if app.Opt.Indexer.Enabled {
		err = indexer.Initialize(ctx, app.Log, indexer.Config{
			Corteza: app.Opt.Corteza,
			ES:      app.Opt.ES,
			Indexer: app.Opt.Indexer,
		})
		if err != nil {
			return
		}
	}

	if app.Opt.Searcher.Enabled {
		err = searcher.Initialize(ctx, app.Log, searcher.Config{
			Corteza:    app.Opt.Corteza,
			ES:         app.Opt.ES,
			HttpServer: app.Opt.HTTPServer,
			Searcher:   app.Opt.Searcher,
		})
		if err != nil {
			return
		}
	}

	app.lvl = bootLevelServicesInitialized
	return
}

// Activate start all internal services and watchers
func (app *CortezaDiscoveryApp) Activate(ctx context.Context) (err error) {
	if app.lvl >= bootLevelActivated {
		return
	}

	if err := app.InitServices(ctx); err != nil {
		return err
	}

	indexer.Watchers(ctx)

	app.lvl = bootLevelActivated

	return nil
}
